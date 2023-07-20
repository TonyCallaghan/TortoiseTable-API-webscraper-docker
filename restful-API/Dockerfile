# Build stage
FROM maven:3.8.1-openjdk-17-slim AS build
WORKDIR /app
COPY . /app
RUN mvn clean package

# Package stage
FROM openjdk:17-alpine
COPY --from=build /app/target/*.jar app.jar
ENTRYPOINT ["java","-jar","/app.jar"]

