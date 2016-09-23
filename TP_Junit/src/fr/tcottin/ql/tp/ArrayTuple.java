package fr.tcottin.ql.tp;

import java.util.ArrayList;
import java.util.List;

public class ArrayTuple{

    private List<Tuple> tupleArray;

    public ArrayTuple(List<Tuple> list){
        if(list.size() != 10){
            throw new ArrayStoreException("Need 10 elements");
        }
        this.tupleArray = new ArrayList<>();
        tupleArray.addAll(list);
    }

    public int getScore(){
        if(tupleArray == null)
            throw new NullPointerException();
        int sum = 0;
        for(int i = 0; i < tupleArray.size(); i++){
            int score = tupleArray.get(i).add();
            sum += score;
            if(tupleArray.get(i).getFirst()==10){ // strike!
                if(i+1 < tupleArray.size()){
                    int val= tupleArray.get(i+1).getFirst();
                    sum += val;
                    if(val == 10){
                        if(i+2 < tupleArray.size())
                            sum += tupleArray.get(i+2).getFirst();
                    }
                    else{
                        sum += tupleArray.get(i+1).getSecond();
                    }
                }
            }
            else if(score == 10){
                if(i+1 < tupleArray.size()){
                    sum += tupleArray.get(i+1).getFirst();
                }
            }
        }
        return sum;
    }

}
