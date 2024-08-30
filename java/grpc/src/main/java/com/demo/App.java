package com.demo;

/**
 * Hello world!
 *
 */
public class App {
    public static void main(String[] args) {
        String greeting = sayHello("John");
        System.out.println(greeting);
    }

    public static String sayHello(String name) {
        return "Hello " + name;
    }
}
