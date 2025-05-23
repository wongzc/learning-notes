package accessmodifier.other;

import accessmodifier.parent.Parent;

public class Unrelated {
    public void access() {
        Parent p = new Parent();
        // System.out.println(p.privateVar);   // ❌ Error: private
        // System.out.println(p.defaultVar);   // ❌ Error: default not visible
        // System.out.println(p.protectedVar); // ❌ Error: protected not accessible from unrelated class
        System.out.println(p.publicVar);       // ✅ OK: public
    }
}
