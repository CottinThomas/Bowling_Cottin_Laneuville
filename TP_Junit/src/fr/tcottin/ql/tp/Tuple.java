package fr.tcottin.ql.tp;

import java.util.List;

/**
 * Created by thomas on 21/09/16.
 */
public class Tuple {

    private int first;
    private int second;

    public Tuple(int first, int second) {
        if(first + second > 10)
            throw new IllegalArgumentException("Sum must be < 10");
        if(first < 0 || second < 0)
            throw new IllegalArgumentException("Values must be positive");
        this.first = first;
        this.second = second;
    }

    public int add(){
        return this.first+this.second;
    }

    public int getFirst() {
        return first;
    }

    public void setFirst(int first) {
        this.first = first;
    }

    public int getSecond() {
        return second;
    }

    public void setSecond(int second) {
        this.second = second;
    }
}
