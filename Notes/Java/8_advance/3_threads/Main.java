//


// create thread using class and overriding run()
// public class Main extends Thread {
//   public static void main(String[] args) {
//     Main thread = new Main();
//     thread.start();
//     System.out.println("This code is outside of the thread");
//   }
//   public void run() {
//     System.out.println("This code is running in a thread");
//   }
// }

// create thread by implement Runnable interface
// public class Main implements Runnable {
//   public static void main(String[] args) {
//     Main obj = new Main();
//     Thread thread = new Thread(obj);
//     thread.start();
//     System.out.println("This code is outside of the thread");
//   }
//   public void run() {
//     System.out.println("This code is running in a thread");
//   }
// }

public class Main extends Thread {
  public static int amount = 0; // share across all instance of Main

  public static void main(String[] args) {
    Main thread = new Main();
    thread.start(); // start a new thread, then it will run the run()

    // use .isAlive to check if a thread finished
    while(thread.isAlive()) {
      System.out.println("Waiting...");
    }
    // Update amount and print its value
    System.out.println("Main: " + amount);
    amount++;
    System.out.println("Main: " + amount);
  }

  public void run() { // method override for Thread class, to define what thread should do when start
    amount++;
  }
}
