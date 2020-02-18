import java.io.BufferedWriter;
import java.io.FileWriter;
import java.io.IOException;

public class MonitorFile extends Thread {

    Garden garden;
    BufferedWriter writer = new BufferedWriter(new FileWriter("output.txt"));

    MonitorFile(Garden garden) throws IOException {
        this.garden = garden;
    }

    @Override
    public void run() {
        while(!Thread.interrupted()) {
            int [][] gardenState = {{}};

            try {
                Thread.sleep(100);
                gardenState = garden.getGarden();
            } catch (InterruptedException e) {
                try {
                    writer.close();
                } catch (IOException ex) {
                    ex.printStackTrace();
                }
                e.printStackTrace();
                Thread.currentThread().interrupt();
            }

            try {
                for (int[] row : gardenState) {
                    for (int area : row) {
                            writer.append(String.valueOf(area)).append(" ");
                    }
                    writer.append("\n");
                }

                writer.append("\n");
            } catch (IOException e) {
                e.printStackTrace();
            }
        }

        try {
            writer.close();
        } catch (IOException e) {
            e.printStackTrace();
        }
    }
}