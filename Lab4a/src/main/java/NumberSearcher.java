import java.util.List;
import java.util.Random;

public class NumberSearcher extends Thread {

    Phonebook phonebook;
    List<String> names;
    Random random = new Random();

    NumberSearcher(Phonebook phonebook, List<String> names) {
        this.phonebook = phonebook;
        this.names = names;
    }

    @Override
    public void run() {
        while(!Thread.interrupted()) {
            int number = -1;
            String name = names.get(random.nextInt(names.size()));

            try {
                number = phonebook.searchNumber(name);
            } catch (InterruptedException e) {
                e.printStackTrace();
            }

            if (number == -1)
                System.out.println("Person " + name + " not found");

            else
                System.out.println(name + " is the owner of number " + number);
        }
    }
}
