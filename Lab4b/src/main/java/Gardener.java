public class Gardener extends Thread {

    Garden garden;

    Gardener(Garden garden) {
        this.garden = garden;
    }

    @Override
    public void run() {
        while(!Thread.interrupted()) {
            try {
                Thread.sleep(30);
                garden.setFirstDeadToAlive();
            } catch (InterruptedException e) {
                e.printStackTrace();
                Thread.currentThread().interrupt();
                continue;
            }
        }
    }
}