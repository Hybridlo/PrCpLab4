public class Gardener extends Thread {

    Garden garden;

    Gardener(Garden garden) {
        this.garden = garden;
    }

    @Override
    public void run() {
        while(!Thread.interrupted()) {
            int [][] gardenState = {{}};
            try {
                Thread.sleep(30);
                gardenState = garden.getGardenAndLock();
            } catch (InterruptedException e) {
                e.printStackTrace();
                Thread.currentThread().interrupt();
            }

            for (int i = 0; i < gardenState.length; i++)
                for (int j = 0; j < gardenState[i].length; j++)
                    if (gardenState[i][j] == 0)
                        garden.setAliveAndUnlock(i, j);

        }
    }
}