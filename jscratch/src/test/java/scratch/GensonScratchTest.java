package scratch;

import static org.junit.Assert.*;
import org.junit.Test;
import com.owlike.genson.JsonBindingException;
import scratch.GensonScratch.*;

public class GensonScratchTest {
    @Test
    public void dataFromStr() {
        final String data =
                " { " +
                "   \"str\": \"test\", " +
                "   \"num\": 100, " +
                "   \"jsonObj\": { \"str\": \"abc\"}" +
                " }";
        final JsonRootObj root = GensonScratch.buildRootFromStr(data);
        assertEquals(root.getNum(), 100);
        assertEquals(root.getStr(), "test");
        assertEquals(root.getJsonObj().getStr(), "abc");
    }

    @Test
    public void strFromData() {
        final JsonRootObj root = new JsonRootObj();
        root.setNum(99);
        root.setStr("rootStr");
        root.setJsonObj(new JsonObj());
        root.getJsonObj().setStr("subStr");
        final String jsonStr = GensonScratch.buildJsonStr(root);
        final JsonRootObj test = GensonScratch.buildRootFromStr(jsonStr);
        assertEquals(root.getNum(), test.getNum());
        assertEquals(root.getStr(), test.getStr());
        assertEquals(root.getJsonObj().getStr(), test.getJsonObj().getStr());
    }

    @Test(expected = JsonBindingException.class)
    public void basJsonStr() {
        GensonScratch.buildRootFromStr("{ \"str\" bad json fmt }");
    }
}
