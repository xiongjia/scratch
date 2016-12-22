package scratch;

import static org.junit.Assert.*;

import java.io.IOException;

import org.junit.Test;

import okhttp3.HttpUrl;
import okhttp3.mockwebserver.MockResponse;
import okhttp3.mockwebserver.MockWebServer;

import org.junit.Before;
import org.junit.After;

public class OkHttpScratchTest {
    private final MockWebServer mockServ = new MockWebServer();

    @Before
    public final void setUp() throws IOException {
        /* disable mock server logger */
        java.util.logging.Logger.getLogger(MockWebServer.class.getName())
            .setLevel(java.util.logging.Level.OFF);
        mockServ.start();
    }

    @After
    public final void tearDown() throws IOException {
        /* stop mock server */
        mockServ.shutdown();
    }

    @Test
    public void baseTest() throws IOException {
        final String mockBody = this.getClass().getName();
        final HttpUrl mockUrl = mockServ.url("/");
        mockServ.enqueue(new MockResponse().setResponseCode(200).setBody(mockBody));

        OkHttpScratch okHttp = new OkHttpScratch();
        OkHttpScratch.ResponseData rep = okHttp.get(mockUrl.toString());

        assertEquals(rep.getCode(), 200);
        assertEquals(rep.getBody(), mockBody);
    }
}
