import java.util.Scanner;

public class PrintFancy {
  public static void main(String[] args) {

    Scanner input = new Scanner(System.in);

    for (int i = 0; i<10; i++) {
      System.out.printf("Loop %d... (enter to continue)\n", i);
      input.nextLine();
      clearScreen();
    }

  }

  public static void clearScreen() {
    System.out.print("\033c");
    System.out.flush();
  }
}
