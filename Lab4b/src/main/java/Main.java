import java.io.IOException;

public class Main {
    public static void main(String[] args) throws IOException, InterruptedException {
        Garden garden = new Garden();

        Gardener gardener = new Gardener(garden);
        Nature nature = new Nature(garden);
        MonitorConsole monitorConsole = new MonitorConsole(garden);
        MonitorFile monitorFile = new MonitorFile(garden);

        gardener.start();
        nature.start();
        monitorConsole.start();
        monitorFile.start();

        Thread.sleep(1000);

        gardener.interrupt();
        nature.interrupt();
        monitorConsole.interrupt();
        monitorFile.interrupt();
    }
}
