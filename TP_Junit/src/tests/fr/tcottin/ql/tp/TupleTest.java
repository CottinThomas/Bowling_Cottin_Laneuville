package tests.fr.tcottin.ql.tp;

import fr.tcottin.ql.tp.Tuple;
import org.junit.Test;

import static org.junit.Assert.*;

/**
 * Created by Thomas COTTIN on 21/09/16.
 */
public class TupleTest {

    @Test
    public void testConstructor(){
        Tuple t = null;
        try{
            t = new Tuple(12,2);
        }catch(IllegalArgumentException e){
            e.printStackTrace();
        }
        assertNull(t);
        try{
            t = new Tuple(-5,2);
        }catch(IllegalArgumentException e){
            e.printStackTrace();
        }
        assertNull(t);
        t = new Tuple(2,5);
        assertNotNull(t);
    }

    @Test
    public void testAdd(){
        Tuple t = new Tuple(0,0);
        assertEquals(t.add(),0);
        t = new Tuple(3,2);
        assertEquals(t.add(),5);
    }

}