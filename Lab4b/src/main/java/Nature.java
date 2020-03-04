import java.util.Random;

public class Nature extends Thread {

    Garden garden;

    Nature(Garden garden) {
        this.garden = garden;
    }

    @Override
    public void run() {
        while(!Thread.interrupted()) {
            try {
                Thread.sleep(10);
                garden.flipRandomState();
            } catch (InterruptedException e) {
                e.printStackTrace();
                Thread.currentThread().interrupt();
                continue;
            }


        }
    }
}