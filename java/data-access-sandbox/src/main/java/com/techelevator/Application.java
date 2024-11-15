package com.techelevator;

// mvn -q package && mvn -q exec:java -Dexec.mainClass=com.techelevator.Application

import java.util.List;
import java.util.ArrayList;

import org.apache.commons.dbcp2.BasicDataSource;
import org.springframework.jdbc.core.JdbcTemplate;
import org.springframework.jdbc.support.rowset.SqlRowSet;

import com.techelevator.model.Book;
import com.techelevator.model.BookMapper;

public class Application {

    public static void main(String[] args) {
        BasicDataSource dataSource = new BasicDataSource();
        dataSource.setUrl("jdbc:postgresql://db:5432/library");
        dataSource.setUsername("postgres");
        dataSource.setPassword("postgres1");

        JdbcTemplate jdbcTemplate = new JdbcTemplate(dataSource);

        // Just querying for one value, use `queryForObject` and tell it what class to expect.
        // In this case we're just getting back a String so we pass `String.class`
        {
          String query = """
            SELECT version()
            """;
          String version = jdbcTemplate.queryForObject(query, String.class);

          System.out.printf("Your database is running %s\n\n", version);
        }


        {
          System.out.println("********************************************************************************");
          System.out.println("* A book:");
          System.out.println("********************************************************************************");

          // To query for something more complicated than a string you have to tell it how to translate the
          // results by giving it a "RowMapper" like our `BookMapper` class.
          String query = """
            SELECT book_id, title, copies
            FROM books
            ORDER BY book_id
            LIMIT 1
          """;
          Book book = jdbcTemplate.queryForObject(query, new BookMapper());
          System.out.printf("%s is \"%s\"\n\n", book.getID(), book.getTitle());
        }


        {
          System.out.println("********************************************************************************");
          System.out.println("* All books:");
          System.out.println("********************************************************************************");


          String query = """
            SELECT book_id, title, copies
            FROM books
          """;
          List<Book> books = jdbcTemplate.query(query, new BookMapper());

          printBooks(books);
        }

        {

          String query = """
            SELECT book_id, title, copies
            FROM books
          """;

          List<Book> books = new ArrayList<>();

          SqlRowSet rows = jdbcTemplate.queryForRowSet(query);
          while (rows.next()) {
            books.add(mapRowToBook(rows));
          }

          printBooks(books);
        }

        {
          System.out.println("********************************************************************************");
          System.out.println("* Some books:");
          System.out.println("********************************************************************************");

          // To pass arguments to your query you write placeholders in your query (the question marks) then
          // pass values in as additional arguments to `query`
          String query = """
            SELECT book_id, title, copies
            FROM books
            WHERE copies >= ?
              AND title ILIKE ?
          """;
          String searchTerm = "%Goose%";
          int copiesMinimum = 1;
          List<Book> books = jdbcTemplate.query(query, new BookMapper(), copiesMinimum, searchTerm);

          printBooks(books);
        }
    }


    public static void printBooks(List<Book> books) {
      int result = 1;
      for (Book book: books) {
        System.out.printf(
            "Result %d:\nid: %s\ntitle: \"%s\"\ncopies: %d\n\n",
            result,
            book.getID(),
            book.getTitle(),
            book.getCopies());
        result++;
      }
    }
}
