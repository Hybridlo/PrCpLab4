import com.sun.tools.javac.util.Pair;

import java.util.List;

public class Main {
    public static void main(String[] args) throws InterruptedException {
        List<Pair<String, Integer>> book = Misc.createPhonebook();

        Phonebook phonebook = new Phonebook();

        Writer writer = new Writer(phonebook, book);
        NameSearcher nameSearcher = new NameSearcher(phonebook, book.size());
        NumberSearcher numberSearcher = new NumberSearcher(phonebook, Misc.getNames(book));

        writer.start();
        nameSearcher.start();
        numberSearcher.start();

        writer.join();
        Thread.sleep(10);
        nameSearcher.interrupt();
        numberSearcher.interrupt();
    }
}
