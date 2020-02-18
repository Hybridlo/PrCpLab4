import com.sun.tools.javac.util.Pair;

import java.util.ArrayList;
import java.util.List;

public class Phonebook {

    List<Pair<String, Integer>> book = new ArrayList<>();

    private int readers = 0;
    private int writers = 0;
    private int writeRequests = 0;

    String searchName(int number) throws InterruptedException {
        lockRead();

        for(Pair<String, Integer> entry : book) {
            if (entry.snd == number) {
                unlockRead();
                return entry.fst;
            }
        }

        unlockRead();
        return "Not found";
    }

    int searchNumber(String name) throws InterruptedException {
        lockRead();

        for(Pair<String, Integer> entry : book) {
            if (entry.fst.equals(name)) {
                unlockRead();
                return entry.snd;
            }
        }

        unlockRead();
        return -1;
    }

    void addEntry(String name, int number) throws InterruptedException {
        lockWrite();

        book.add(new Pair<>(name, number));

        unlockWrite();
    }

    private synchronized void lockRead() throws InterruptedException{
        while(writers > 0 || writeRequests > 0){
            wait();
        }
        readers++;
    }

    private synchronized void unlockRead() {
        readers--;
        notifyAll();
    }

    private synchronized void lockWrite() throws InterruptedException{
        writeRequests++;

        while(readers > 0 || writers > 0){
            wait();
        }
        writeRequests--;
        writers++;
    }

    private synchronized void unlockWrite() {
        writers--;
        notifyAll();
    }
}
