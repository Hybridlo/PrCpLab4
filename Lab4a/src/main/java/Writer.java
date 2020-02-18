import com.sun.tools.javac.util.Pair;

import java.util.List;

public class Writer extends Thread {

    Phonebook phonebook;
    List<Pair<String, Integer>> book;

    Writer(Phonebook phonebook, List<Pair<String, Integer>> book) {
        this.phonebook = phonebook;
        this.book = book;
    }

    @Override
    public void run() {
        while(!Thread.interrupted()) {
            for (Pair<String, Integer> entry : book) {

                try {
                    phonebook.addEntry(entry.fst, entry.snd);
                    System.out.println("Wrote " + entry.fst + " with number " + entry.snd);
                } catch (InterruptedException e) {
                    e.printStackTrace();
                }

            }

            break;
        }
    }
}
