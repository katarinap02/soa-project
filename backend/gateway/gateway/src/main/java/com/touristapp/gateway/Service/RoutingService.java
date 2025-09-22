package com.touristapp.gateway.Service;

import com.fasterxml.jackson.databind.JsonNode;
import com.fasterxml.jackson.databind.ObjectMapper;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.*;
import org.springframework.stereotype.Service;
import org.springframework.web.client.RestTemplate;
import jakarta.servlet.http.HttpServletRequest;

import java.util.Collections;
import java.util.HashMap;
import java.util.Map;

@Service
public class RoutingService {
    
    @Autowired
    private RestTemplate restTemplate;
    
    @Autowired
    private StakeholderRpcClient stakeholderRpcClient;
    
    private final ObjectMapper objectMapper = new ObjectMapper();
    
    private final Map<String, String> serviceRoutes = new HashMap<String, String>() {{
        put("stakeholders", "http://stakeholder_server:8080");
        put("blog", "http://blog_server:8081");
        put("tours", "http://tours_server:8082");
        put("followers", "http://followers_server:8084");
    }};
    
    public ResponseEntity<?> forwardRequest(HttpServletRequest request, String body) {
        String serviceName = determineTargetService(request.getRequestURI());
        String fullPath = request.getRequestURI();
        String method = request.getMethod();
        
        // Check if this specific request should use RPC
        if (shouldUseRpc(serviceName, fullPath, method)) {
            return handleRpcRequest(fullPath, method, request, body);
        }
        
        // Otherwise use regular HTTP forwarding (your existing logic)
        return handleHttpRequest(request, body);
    }
    
    /**
     * Determines if a specific request should be routed via RPC
     */
    private boolean shouldUseRpc(String serviceName, String fullPath, String method) {
        // Only these 2 specific stakeholder endpoints use RPC
        return "stakeholders".equals(serviceName) && (
            ("GET".equals(method) && fullPath.equals("/stakeholders/users")) ||
            ("POST".equals(method) && fullPath.equals("/stakeholders/users/login"))
        );
    }
    
    /**
     * Handles RPC routing for specific endpoints
     */
    private ResponseEntity<?> handleRpcRequest(String fullPath, String method, 
                                             HttpServletRequest request, String body) {
        try {
            String result;
            
            if ("/stakeholders/users".equals(fullPath) && "GET".equals(method)) {
                // RPC call for getting all users
                result = stakeholderRpcClient.getAllUsers();
                
            } else if ("/stakeholders/users/login".equals(fullPath) && "POST".equals(method)) {
                // Parse login credentials from request body
                JsonNode loginData = objectMapper.readTree(body);
                String username = loginData.get("username").asText();
                String password = loginData.get("password").asText();
                
                // RPC call for user login
                result = stakeholderRpcClient.loginUser(username, password);
                
            } else {
                return ResponseEntity.badRequest()
                    .body("{\"error\":\"RPC endpoint not implemented\"}");
            }
            
            // Return successful RPC response
            return ResponseEntity.ok()
                .header("Content-Type", "application/json")
                .header("X-Gateway-Method", "RPC")
                .header("X-Service", "stakeholder")
                .body(result);
                
        } catch (Exception e) {
            // RPC call failed, return error
            return ResponseEntity.status(HttpStatus.SERVICE_UNAVAILABLE)
                .header("X-Gateway-Method", "RPC")
                .body("{\"error\":\"RPC call failed: " + e.getMessage() + "\"}");
        }
    }
    
    /**
     * Handles traditional HTTP routing (your existing logic)
     */
    private ResponseEntity<?> handleHttpRequest(HttpServletRequest request, String body) {
        String serviceName = determineTargetService(request.getRequestURI());
        String remainingPath = extractRemainingPath(request.getRequestURI(), serviceName, request);
        String targetUrl = serviceRoutes.get(serviceName) + remainingPath;
        
        // Add query parameters if present
       // if (request.getQueryString() != null) {
          //  targetUrl += "?" + request.getQueryString();
       // }
        
        // Copy headers from original request
        HttpHeaders headers = new HttpHeaders();
        Collections.list(request.getHeaderNames()).forEach(headerName ->
            headers.set(headerName, request.getHeader(headerName))
        );
        
        if (!headers.containsKey("Content-Type")) {
            headers.setContentType(MediaType.APPLICATION_JSON);
        }

        // Add gateway identification headers
        headers.set("X-Gateway-Method", "HTTP");
        headers.set("X-Gateway-Source", "tourist-gateway");
        
        HttpEntity<String> entity = new HttpEntity<>(body, headers);
        
        try {
            return restTemplate.exchange(
                targetUrl,
                HttpMethod.valueOf(request.getMethod()),
                entity,
                String.class
            );
        } catch (Exception e) {
            return ResponseEntity.status(HttpStatus.SERVICE_UNAVAILABLE)
                .header("X-Gateway-Method", "HTTP")
                .body("{\"error\":\"HTTP service unavailable: " + e.getMessage() + "\"}");
        }
    }
    
    private String determineTargetService(String uri) {
        String[] pathParts = uri.split("/");
        if (pathParts.length >= 2) {
            return pathParts[1];
        }
        return "default";
    }
    
    private String extractRemainingPath(String uri, String serviceName, HttpServletRequest request) {
    String prefix = "/" + serviceName;
    String path = uri.startsWith(prefix) ? uri.substring(prefix.length()) : uri;
    String query = request.getQueryString();
    if (query != null && !query.isEmpty()) {
        path += "?" + query;
    }
    return path;
    }
}

