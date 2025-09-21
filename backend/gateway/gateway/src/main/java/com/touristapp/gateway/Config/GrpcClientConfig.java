package com.touristapp.gateway.Config;

import com.touristapp.gateway.grpc.StakeholderServiceGrpc;
import io.grpc.ManagedChannel;
import io.grpc.ManagedChannelBuilder;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

@Configuration
public class GrpcClientConfig {

    @Bean
    public ManagedChannel stakeholderChannel() {
        return ManagedChannelBuilder.forAddress("stakeholders", 9091)
                .usePlaintext()
                .keepAliveWithoutCalls(true)
                .build();
    }

    @Bean
    public StakeholderServiceGrpc.StakeholderServiceBlockingStub stakeholderServiceStub(
            ManagedChannel stakeholderChannel) {
        return StakeholderServiceGrpc.newBlockingStub(stakeholderChannel);
    }
}
