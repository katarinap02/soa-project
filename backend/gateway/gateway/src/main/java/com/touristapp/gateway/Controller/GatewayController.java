package com.touristapp.gateway.Controller;

import com.touristapp.gateway.Service.AuthService;
import com.touristapp.gateway.Service.RoutingService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import jakarta.servlet.http.HttpServletRequest;

@RestController
public class GatewayController {

    @Autowired
    private AuthService authService;

    @Autowired
    private RoutingService routingService;

    @RequestMapping(value = "/**", method = {RequestMethod.GET, RequestMethod.POST, RequestMethod.PUT, RequestMethod.DELETE})
    public ResponseEntity<?> handleRequest(
            HttpServletRequest request,
            @RequestBody(required = false) String body) {

        if (!authService.isAuthorized(request)) {
            return ResponseEntity.status(HttpStatus.UNAUTHORIZED)
                    .body("{\"error\":\"Unauthorized\"}");
        }

        return routingService.forwardRequest(request, body);
    }
}
