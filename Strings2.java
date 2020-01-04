import java.util.HashMap;
import java.util.ArrayList;
import java.util.List;


class Helper {
    HashMap<Character,Integer> hMap;
    
    Boolean areAnagrams(String first, String second) {

        hMap = new HashMap<Character,Integer>();
        if (first.length() != second.length()) {
            return false;
        }
        for (int i = 0; i < first.length(); i++) {
            char c = first.charAt(i);
            if (hMap.containsKey(c)) {
                hMap.replace(c, hMap.get(c)+1);
            } else {
                hMap.put(c, 1);
            }
        }
        for (int i = 0; i < second.length(); i++) {
            char c = second.charAt(i);
            if (hMap.containsKey(c) && hMap.get(c) != 0) {
                hMap.replace(c, hMap.get(c)-1);
            } else {
                return false;
            }
        }
        for (char c : hMap.keySet()) {
            if (hMap.get(c) != 0) {
                return false;
            }
        }
        return true;
    }
}

public class Strings2 {
    public static void main(String[] args) {
        Helper helper = new Helper();
        List<String> words1 = new ArrayList<String>();
        List<String> words2 = new ArrayList<String>();
        words1.add("kayak");
        words2.add("ayakk");
        words1.add("probowwow");
        words2.add("notprobowwow");
        words1.add("kty");
        words2.add("mtw");
        for (int i = 0; i < words1.size(); i++) {
            Boolean result = helper.areAnagrams(words1.get(i), words2.get(i));
            System.out.println("Result is " + result + " for " + words1.get(i) + " and " + words2.get(i));
        }
        return;
    }
}