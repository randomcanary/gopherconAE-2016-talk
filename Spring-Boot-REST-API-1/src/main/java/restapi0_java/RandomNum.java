package restapi0_java;

import java.util.concurrent.ThreadLocalRandom;

public class RandomNum {
	
	private final int RandomInt;
	
	public RandomNum() {
		this.RandomInt = ThreadLocalRandom.current().nextInt(0, 101);
	}
	
	public int getRandomInt() {
		return RandomInt;
	}
}

