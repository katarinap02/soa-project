package com.touristapp.gateway.Service;

import org.springframework.stereotype.Service;
import jakarta.servlet.http.HttpServletRequest;

@Service
public class AuthService {

    public boolean isAuthorized(HttpServletRequest request) {
        /*
        String authHeader = request.getHeader("Authorization");
        
        if (authHeader == null || !authHeader.startsWith("Bearer ")) {
            return false;
        }

        String token = authHeader.substring(7);
        return validateToken(token);
        */
       return true;    
    }

    private boolean validateToken(String token) {
        return token != null && !token.trim().isEmpty();
    }
}