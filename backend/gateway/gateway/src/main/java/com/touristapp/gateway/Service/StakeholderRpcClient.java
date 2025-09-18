package com.touristapp.gateway.Service;

import com.fasterxml.jackson.databind.ObjectMapper;
import com.touristapp.gateway.grpc.*;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.stream.Collectors;

@Service
public class StakeholderRpcClient {

    @Autowired
    private StakeholderServiceGrpc.StakeholderServiceBlockingStub stakeholderServiceStub;

    private final ObjectMapper objectMapper = new ObjectMapper();

    public String getAllUsers() throws Exception {
        try {
            GetAllUsersRequest request = GetAllUsersRequest.newBuilder().build();
            GetAllUsersResponse response = stakeholderServiceStub.getAllUsers(request);
            
            // Convert protobuf response to JSON
            Map<String, Object> result = new HashMap<>();
            result.put("success", response.getSuccess());
            result.put("message", response.getMessage());
            
            List<Map<String, Object>> users = response.getUsersList().stream()
                .map(this::convertUserToMap)
                .collect(Collectors.toList());
            result.put("users", users);
            
            return objectMapper.writeValueAsString(result);
            
        } catch (Exception e) {
            throw new Exception("Failed to get all users via RPC: " + e.getMessage());
        }
    }

    public String loginUser(String email, String password) throws Exception {
        try {
            LoginUserRequest request = LoginUserRequest.newBuilder()
                .setEmail(email)
                .setPassword(password)
                .build();
                
            LoginUserResponse response = stakeholderServiceStub.loginUser(request);
            
            // Convert protobuf response to JSON
            Map<String, Object> result = new HashMap<>();
            result.put("success", response.getSuccess());
            result.put("message", response.getMessage());
            
            if (response.getSuccess()) {
                result.put("token", response.getToken());
                result.put("user", convertUserToMap(response.getUser()));
            }
            
            return objectMapper.writeValueAsString(result);
            
        } catch (Exception e) {
            throw new Exception("Failed to login user via RPC: " + e.getMessage());
        }
    }

    private Map<String, Object> convertUserToMap(User user) {
        Map<String, Object> userMap = new HashMap<>();
        userMap.put("id", user.getId());
        userMap.put("email", user.getEmail());
        userMap.put("name", user.getName());
        userMap.put("role", user.getRole());
        userMap.put("created_at", user.getCreatedAt());
        return userMap;
    }
}