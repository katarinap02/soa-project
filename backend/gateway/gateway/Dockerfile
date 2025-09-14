FROM openjdk:24-jdk-slim

WORKDIR /app

# Copy gradle wrapper and build files
COPY gradlew .
COPY gradle gradle
COPY build.gradle .
COPY settings.gradle .

# Download dependencies
RUN chmod +x ./gradlew && ./gradlew dependencies

# Copy source code
COPY src ./src

# Build the application
RUN ./gradlew clean build -x test

# Use the Spring Boot fat JAR (without -plain suffix)
EXPOSE 8085
ENTRYPOINT ["java", "-jar", "build/libs/gateway-0.0.1-SNAPSHOT.jar"]
