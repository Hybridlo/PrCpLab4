import com.sun.tools.javac.util.Pair;

import java.util.ArrayList;
import java.util.List;

public class Misc {
    private Misc(){};

    public static final ArrayList<Pair<String, Integer>> createPhonebook() {
        ArrayList<Pair<String, Integer>> result = new ArrayList<>();
        result.add(new Pair<>("Fifteen", 15));
        result.add(new Pair<>("Four", 4));
        result.add(new Pair<>("Nine", 9));
        result.add(new Pair<>("Three", 3));
        result.add(new Pair<>("Thirteen", 13));
        result.add(new Pair<>("Eighteen", 18));
        result.add(new Pair<>("Seven", 7));
        result.add(new Pair<>("One", 1));
        result.add(new Pair<>("Seventeen", 17));
        result.add(new Pair<>("Twelve", 12));
        result.add(new Pair<>("Six", 6));
        result.add(new Pair<>("Zero", 0));
        result.add(new Pair<>("Nineteen", 19));
        result.add(new Pair<>("Eleven", 11));
        result.add(new Pair<>("Two", 2));
        result.add(new Pair<>("Five", 5));
        result.add(new Pair<>("Eight", 8));
        result.add(new Pair<>("Fourteen", 14));
        result.add(new Pair<>("Ten", 10));
        result.add(new Pair<>("Sixteen", 16));

        return result;
    }

    public static final ArrayList<String> getNames(List<Pair<String, Integer>> book) {
        ArrayList<String> names = new ArrayList<>();

        for (Pair<String, Integer> entry : book) {
            names.add(entry.fst);
        }

        return names;
    }
}
