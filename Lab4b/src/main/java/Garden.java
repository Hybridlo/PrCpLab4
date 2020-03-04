import java.util.Random;
import java.util.concurrent.locks.ReadWriteLock;
import java.util.concurrent.locks.ReentrantReadWriteLock;

public class Garden {

    int[][] garden = {
            {1, 1, 1, 1, 1},
            {1, 1, 1, 1, 1},
            {1, 1, 1, 1, 1},
            {1, 1, 1, 1, 1},
            {1, 1, 1, 1, 1}
    };

    private ReadWriteLock lock = new ReentrantReadWriteLock();

    void setFirstDeadToAlive() {
        lock.writeLock().lock();

        boolean changed = false;

        for (int i = 0; i < garden.length; i++) {
            for (int j = 0; j < garden[i].length; j++) {
                if (garden[i][j] == 0) {
                    garden[i][j] = 1;
                    changed = true;
                    break;
                }
            }

            if (changed)
                break;
        }

        lock.writeLock().unlock();
    }

    void flipRandomState() {
        lock.writeLock().lock();

        Random random = new Random();

        int i = random.nextInt(garden.length);
        int j = random.nextInt(garden[i].length);

        garden[i][j] = 1 - garden[i][j];

        lock.writeLock().unlock();
    }

    int[][] getGarden() {
        lock.readLock().lock();

        int[][] result = garden;

        lock.readLock().unlock();

        return result;
    }
}
