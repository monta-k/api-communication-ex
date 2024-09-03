package com.demo.server;

import java.io.IOException;

import io.grpc.Server;
import io.grpc.ServerBuilder;

public class CarParkServer {
  public static void main(String[] args) throws IOException, InterruptedException {
    Server server = ServerBuilder.forPort(5003).addService(new CarParkServiceImpl()).build();

    server.start();

    server.awaitTermination();
  }
}
