package com.demo.client;

import com.demo.CarParkServiceGrpc;
import com.demo.ParkRequest;
import com.demo.ParkResponse;
import com.demo.Vehicle;

import io.grpc.ManagedChannel;
import io.grpc.ManagedChannelBuilder;

public class CarParkClient {
  public static void main(String[] args) {
    ManagedChannel channel = ManagedChannelBuilder.forAddress("localhost", 5003).usePlaintext().build();

  CarParkServiceGrpc.CarParkServiceBlockingStub stub = CarParkServiceGrpc.newBlockingStub(channel);

  ParkRequest parkRequest = ParkRequest.newBuilder().setVehicle(Vehicle.newBuilder().setVehicleNumber("NA-1324").setVehicleType("BUS").build()).build();

  ParkResponse parkResponse = stub.parkVehicle(parkRequest);
  
  System.out.println("Response for the first call: " + parkResponse.getResult());
  
  }
}
