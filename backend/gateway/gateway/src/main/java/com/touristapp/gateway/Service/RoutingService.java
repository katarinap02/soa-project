package com.touristapp.gateway.Service;

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
    
    private final Map<String, String> serviceRoutes = new HashMap<String, String>() {{
        put("stakeholders", "http://stakeholder_server:8080");
        put("blog", "http://blog_server:8081");
        put("tours", "http://tours_server:8082");
        put("followers", "http://followers_server:8084");
    }};

    public ResponseEntity<?> forwardRequest(HttpServletRequest request, String body) {
        String serviceName = determineTargetService(request.getRequestURI());
        String remainingPath = extractRemainingPath(request.getRequestURI(), serviceName, request);
        String targetUrl = serviceRoutes.get(serviceName) + remainingPath;
        
        HttpHeaders headers = new HttpHeaders();
        Collections.list(request.getHeaderNames()).forEach(headerName ->
            headers.set(headerName, request.getHeader(headerName))
        );
         if (!headers.containsKey("Content-Type")) {
            headers.setContentType(MediaType.APPLICATION_JSON);
        }
    
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
                .body("{\"error\":\"Service unavailable\"}");
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