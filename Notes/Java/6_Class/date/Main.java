import java.time.LocalDate;
import java.time.LocalDateTime;
import java.time.format.DateTimeFormatter;

public class Main {
  public static void main(String[] args) {
    LocalDate myObj = LocalDate.now(); // use .now() for current time
    System.out.println(myObj); 
    // LocalDate	date yyyy-MM-dd
    // LocalTime	time HH-mm-ss-ns
    // LocalDateTime	date + time yyyy-MM-dd-HH-mm-ss-ns
    // DateTimeFormatter	Formatter for displaying and parsing date-time objects


    //parse date time object
    LocalDateTime myDateObj = LocalDateTime.now();
    System.out.println("Before formatting: " + myDateObj);
    DateTimeFormatter myFormatObj = DateTimeFormatter.ofPattern("dd-MM-yyyy HH:mm:ss");

    String formattedDate = myDateObj.format(myFormatObj);
    System.out.println("After formatting: " + formattedDate);


  }
}