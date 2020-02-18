import java.util.Random;

public class NameSearcher extends Thread {

    Phonebook phonebook;
    int length;
    Random random = new Random();

    NameSearcher(Phonebook phonebook, int length) {
        this.phonebook = phonebook;
        this.length = length;
    }

    @Override
    public void run() {
        while(!Thread.interrupted()) {
            String name = "";
            int number = random.nextInt(length);

            try {
                name = phonebook.searchName(number);
            } catch (InterruptedException e) {
                e.printStackTrace();
            }

            if (name.equals("Not found"))
                System.out.println("Number " + number + " not found");

            else
                System.out.println("Owner of " + number + " is " + name);
        }
    }
}
