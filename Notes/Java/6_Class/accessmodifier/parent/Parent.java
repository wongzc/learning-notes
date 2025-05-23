package accessmodifier.parent;

public class Parent {
    private String privateVar = "private";
    String defaultVar = "default";            // package-private
    protected String protectedVar = "protected";
    public String publicVar = "public";

    public void printAll() {
        System.out.println(privateVar);
        System.out.println(defaultVar);
        System.out.println(protectedVar);
        System.out.println(publicVar);
    }
}