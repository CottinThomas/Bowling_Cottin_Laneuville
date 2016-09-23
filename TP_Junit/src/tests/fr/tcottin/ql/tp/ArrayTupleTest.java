package tests.fr.tcottin.ql.tp;

import fr.tcottin.ql.tp.ArrayTuple;
import fr.tcottin.ql.tp.Tuple;
import org.junit.Before;
import org.junit.Test;

import java.util.ArrayList;
import java.util.List;

import static org.junit.Assert.*;

/**
 * Created by Thomas COTTIN on 21/09/16.
 */
public class ArrayTupleTest {

    public ArrayTuple arrayTuple;

    @Before
    public void init(){

    }


    @Test
    public void testStrikes(){

        List<Tuple> l = new ArrayList<>();
        l.add(new Tuple(10,0));
        l.add(new Tuple(6,1));
        l.add(new Tuple(3,1));
        l.add(new Tuple(1,1));
        l.add(new Tuple(1,1));
        l.add(new Tuple(1,1));
        l.add(new Tuple(1,1));
        l.add(new Tuple(1,1));
        l.add(new Tuple(1,1));
        l.add(new Tuple(1,1));
        arrayTuple = new ArrayTuple(l);

        assertEquals(arrayTuple.getScore(),42);
    }


    @Test
    public void testSpare(){
        List<Tuple> l = new ArrayList<>();
        l.add(new Tuple(8,2));
        l.add(new Tuple(6,1));
        l.add(new Tuple(3,1));
        l.add(new Tuple(1,1));
        l.add(new Tuple(1,1));
        l.add(new Tuple(1,1));
        l.add(new Tuple(1,1));
        l.add(new Tuple(1,1));
        l.add(new Tuple(1,1));
        l.add(new Tuple(1,1));
        arrayTuple = new ArrayTuple(l);

        assertEquals(arrayTuple.getScore(),41);
    }

    @Test
    public void testMultipleSpecials(){
         List<Tuple> l = new ArrayList<>();
        l.add(new Tuple(8,2));
        l.add(new Tuple(10,0));
        l.add(new Tuple(10,0));
        l.add(new Tuple(9,1));
        l.add(new Tuple(1,1));
        l.add(new Tuple(1,1));
        l.add(new Tuple(1,1));
        l.add(new Tuple(1,1));
        l.add(new Tuple(1,1));
        l.add(new Tuple(1,1));
        arrayTuple = new ArrayTuple(l);

        assertEquals(arrayTuple.getScore(),92);
    }
}
