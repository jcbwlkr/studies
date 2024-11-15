public class Switch {
  public static void main(String[] args) {

    int age = 22;
    String country = "USA";

    if ((age >= 21 && country == "USA") ||
        (age > 15 && country == "France")) {
      System.out.println("OK to drink");
    } else {
      System.out.println("Not ok to drink");
    }
  }
}
