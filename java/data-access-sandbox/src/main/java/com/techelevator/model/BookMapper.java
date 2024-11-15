package com.techelevator.model;

import org.springframework.jdbc.core.RowMapper;

import java.sql.ResultSet;
import java.sql.SQLException;

import com.techelevator.model.Book;

public class BookMapper implements RowMapper<Book> {
  public Book mapRow(ResultSet rs, int rowNum) throws SQLException {
    Book book = new Book();
    book.setID(rs.getString("book_id"));
    book.setTitle(rs.getString("title"));
    book.setCopies(rs.getInt("copies"));
    return book;
  }
}





