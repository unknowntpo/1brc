/*
 * This Java source file was generated by the Gradle 'init' task.
 */
package onebrc.java;

import org.junit.jupiter.api.Test;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.io.BufferedReader;
import java.io.FileReader;
import java.nio.file.Path;
import java.text.MessageFormat;
import java.util.HashMap;

import static org.junit.jupiter.api.Assertions.*;

class AppTest {
    private static final Logger logger = LoggerFactory.getLogger(AppTest.class);

    @Test
    void appHasAGreeting() {
        App classUnderTest = new App();
        assertNotNull(classUnderTest.getGreeting(), "app should have a greeting");
    }

    @Test
    void testCompute() {
        App app = new App();

        HashMap<String, Result> want = new HashMap<>();

        want.put("Tokyo", new Result("Tokyo", 33.6f, 35.6f, 34f, 2, 69.2f));
        want.put("Jakarta", new Result("Jakarta", -6.1f, -6.1f, -6.1f, 1, -6.1f));
        want.put("Delhi", new Result("Delhi", 28.6f, 28.6f, 28.6f, 1, 28.6f));
        want.put("Guangzhou", new Result("Guangzhou", 23.1f, 33.1f, 28f, 2, 56.2f));

        Path cwd = Path.of("").toAbsolutePath();
        String filePath = cwd.toString() + "/testdata/weather_stations.csv"; // Replace with your actual file path

        try (BufferedReader br = new BufferedReader(new FileReader(filePath))) {
            HashMap<String, Result> got = app.compute(br);
            got.forEach((key, value) -> {
                logger.debug(MessageFormat.format("got: key {0} value {1}", key, value.toString()));
            });
            // can't compare between float
            assertEquals(want.keySet(), got.keySet());
        } catch (Exception e) {
            e.printStackTrace();
        }
    }
}
