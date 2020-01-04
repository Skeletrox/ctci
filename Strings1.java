import java.util.ArrayList;
import java.util.Arrays;
import java.util.HashSet;
import java.util.List;



class Helper {

    HashSet<Character> wordSet;

    Helper() {
       // pass
    }

    boolean isUniqueChars(String s) {
        wordSet = new HashSet<Character>();
        for (int i = 0; i < s.length(); i++) {
            if (wordSet.contains(s.charAt(i))) {
                return false;
            }
            wordSet.add(s.charAt(i));
        }
        return true;
    }
}

class HelperBonus {

    boolean isUniqueChars(String s) {
        char tempArray[] = s.toCharArray();
        Arrays.sort(tempArray);
        for (int i = 0; i < tempArray.length - 1; i++) {
            if (tempArray[i] == tempArray[i+1]) {
                return false;
            }
        }
        return true;
    }
}

public class Strings1 {

    public static void main(String[] args) {
        Helper helper = new Helper();
        HelperBonus helperBonus = new HelperBonus();
        List<String> words = new ArrayList<String>();
        words.add("abcdef");
        words.add("aaaaaa");
        words.add("abcdefghijklmnopqrstuvwxyza");
        for (String w : words) {
            Boolean res = helper.isUniqueChars(w);
            System.out.println("Uniqueness: " + w +  " is " +  res.toString());
        }
        System.out.println("Bonus\n");
        for (String w : words) {
            Boolean res = helperBonus.isUniqueChars(w);
            System.out.println("Uniqueness: " + w +  " is " +  res.toString());
        }
    }
}