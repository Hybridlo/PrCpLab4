import java.util.ArrayList;
import java.util.List;

public class Garden {

    int[][] garden = {
            {1, 1, 1, 1, 1},
            {1, 1, 1, 1, 1},
            {1, 1, 1, 1, 1},
            {1, 1, 1, 1, 1},
            {1, 1, 1, 1, 1}
    };

    private int readers = 0;
    private int writers = 0;
    private int writeRequests = 0;

    int[][] getGardenAndLock() throws InterruptedException {
        lockWrite();

        return garden;
    }

    void setStateAndUnlock(int row, int column, int state) {

        garden[row][column] = state;

        unlockWrite();
    }

    void setAliveAndUnlock(int row, int column)  {

        garden[row][column] = 1;

        unlockWrite();
    }

    int[][] getGarden() throws InterruptedException {
        lockRead();

        int[][] result = garden;

        unlockRead();

        return result;
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
