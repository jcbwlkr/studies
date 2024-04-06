package com.jcbwlkr.studies;

import org.json.simple.JSONObject;
import org.json.simple.JSONValue;

import java.io.IOException;
import java.net.URI;
import java.net.http.HttpClient;
import java.net.http.HttpRequest;
import java.net.http.HttpResponse;
import java.net.http.HttpResponse.BodyHandlers;

public class App {
  public static void main(String args[]) {
    try {
      HttpClient client = HttpClient.newHttpClient();
      HttpRequest request = HttpRequest.newBuilder()
        .uri(URI.create("https://www.deckofcardsapi.com/api/deck/new/shuffle/?deck_count=1"))
        .build();

      HttpResponse<String> response = client.send(request, BodyHandlers.ofString());

      System.out.printf("Status code: %d\n", response.statusCode());

      String body = response.body();
      System.out.printf("Raw body: %s\n", body);

      JSONObject obj = (JSONObject) JSONValue.parse(body);
      String deckID = (String) obj.get("deck_id");
      System.out.printf("Deck ID: %s\n", deckID);

    } catch(Exception e) {
      System.out.println(e);
    }
  }
}
