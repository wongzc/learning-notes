package accessmodifier.child;

import accessmodifier.parent.Parent;

public class Child extends Parent {
    public void accessFromSubclass() {
        // System.out.println(privateVar);     // ❌ Error: private
        // System.out.println(defaultVar);     // ❌ Error: default not visible across package
        System.out.println(protectedVar);      // ✅ OK: protected accessible in subclass
        System.out.println(publicVar);         // ✅ OK: public accessible everywhere
    }
}
