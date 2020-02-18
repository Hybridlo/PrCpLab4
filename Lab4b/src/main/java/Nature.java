import java.util.Random;

public class Nature extends Thread {

    Garden garden;
    Random random = new Random();

    Nature(Garden garden) {
        this.garden = garden;
    }

    @Override
    public void run() {
        while(!Thread.interrupted()) {
            int [][] gardenState = {{}};
            try {
                Thread.sleep(10);
                gardenState = garden.getGardenAndLock();
            } catch (InterruptedException e) {
                e.printStackTrace();
                Thread.currentThread().interrupt();
            }

            int i = random.nextInt(gardenState.length);
            int j = random.nextInt(gardenState[i].length);

            garden.setStateAndUnlock(i, j, 1 - gardenState[i][j]);

        }
    }
}