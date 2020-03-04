public class MonitorConsole extends Thread {

    Garden garden;

    MonitorConsole(Garden garden) {
        this.garden = garden;
    }

    @Override
    public void run() {
        while(!Thread.interrupted()) {
            int [][] gardenState;

            try {
                Thread.sleep(100);
                gardenState = garden.getGarden();
            } catch (InterruptedException e) {
                e.printStackTrace();
                Thread.currentThread().interrupt();
                continue;
            }

            for (int[] row : gardenState) {
                for (int area : row)
                    System.out.print(area + " ");
                System.out.println();
            }

            System.out.println();
        }
    }
}