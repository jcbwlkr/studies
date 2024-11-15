package com.techelevator.model;


public class Book {
  private String id;
  private String title;
  private int copies;

  public String getID() {
    return this.id;
  }

  public String getTitle() {
    return this.title;
  }

  public int getCopies() {
    return this.copies;
  }

  public void setID(String id) {
    this.id = id;
  }

  public void setTitle(String title) {
    this.title = title;
  }

  public void setCopies(int copies) {
    this.copies = copies;
  }
}
