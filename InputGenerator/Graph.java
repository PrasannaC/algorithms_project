import java.io.BufferedWriter;
import java.io.File;
import java.io.FileWriter;
import java.io.IOException;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Scanner;

public class Graph {

	static HashMap<Integer,List<EdgeAndCost>> graphMap= new HashMap<Integer,List<EdgeAndCost>>();
	public Graph(){
		graphMap = new HashMap<Integer, List<EdgeAndCost>>();
	}
	public static void main(String[] args) {
		// TODO Auto-generated method stub
		Scanner sc =new Scanner(System.in);
		System.out.println("Enter the number of nodes");
        int nodes = sc.nextInt();
        
        System.out.println("Enter the number of edges");
        int edges = sc.nextInt();
        
        System.out.println("Please wait while the graph is being generated..");
        
        //Graph g = new Graph();
        //List<Integer> nodeList = new ArrayList<Integer>(nodes);
        
        for(Integer i=0; i< nodes - 1 && edges > 0 ;i++)
        {
        	for(Integer j=i+1; j< nodes && edges > 0;j++)
        	{
        		int costOfEdge = (int)Math.round(Math.random() * (100 -0));
        		addEdge(i.intValue(),j.intValue(),costOfEdge);
        		edges--;
        	}
        } 
        printGraph();
	}
    public static void addEdge(int source, int destination, int cost)
    {
    	EdgeAndCost ed = new EdgeAndCost(destination, cost);
        if(graphMap.containsKey(source))
        {
        	List<EdgeAndCost> l = graphMap.get(source);
        	l.add(ed);
        }
        else{
        	List<EdgeAndCost> l = new ArrayList<EdgeAndCost>();
        	l.add(ed);
    	    graphMap.put(source, l);
        }
     }
    
    public static void printGraph(){
    	 
    	File file= new File("F:/New Eclipse/my Workspaceruleengine/DAA Project/src/GraphData.txt");
    	FileWriter fr=null;
    	BufferedWriter br=null;
    	try{
    		 fr= new FileWriter(file);
    		 br= new BufferedWriter(fr);
    		
    	
    	for(Integer source:graphMap.keySet())
    	{
    		String  s=source.toString()+" ";
    		System.out.println(s);
    		List<EdgeAndCost> l=graphMap.get(source);
    		for(EdgeAndCost e: l)
    		{
    			s=source.toString()+" "+e.getEdge()+" "+e.getCost()+System.getProperty("line.separator");
    			br.write(s);
    		}
    		
    		
    	}
    /*	Iterator it = graphMap.entrySet().iterator();
    	while (it.hasNext()) {
    		Map.Entry pair = (Map.Entry)it.next();
    	    System.out.println(pair.getKey() + " = " + pair.getValue());
    	    it.remove(); // avoids a ConcurrentModificationException
    			    } */
    	}
    	catch(IOException e)
    	{
    		e.printStackTrace();
    	}
    	finally{
    	try{	
    		br.close();
    		fr.close();
    	}catch(IOException e){
    		e.printStackTrace();
    	}
    }
    }
}
