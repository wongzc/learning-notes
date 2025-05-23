class OuterClass {
  int x = 10;

  class InnerClass {
    int y = 5;
  }

  private class InnerClass2 { // this is private class, only access within same class
    int z = 5;
  }

  // if set to static for inner class, means can acces without create object of outer class
  // but still need to create objet of inner
  static class InnerClass3 {
    int a = 5;
  }

  // inner class ca
}

public class Main {
  public static void main(String[] args) {
    OuterClass myOuter = new OuterClass();
    OuterClass.InnerClass myInner = myOuter.new InnerClass();
    System.out.println(myInner.y + myOuter.x);


    OuterClass.InnerClass3 abc = new OuterClass.InnerClass3();
    System.out.println(abc.a);
  }
}