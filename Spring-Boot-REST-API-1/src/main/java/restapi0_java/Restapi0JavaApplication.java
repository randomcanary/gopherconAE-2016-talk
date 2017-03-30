package restapi0_java;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.EnableAutoConfiguration;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.context.annotation.ComponentScan;
import org.springframework.context.annotation.Configuration;
import org.springframework.web.bind.annotation.RestController;

@SpringBootApplication
public class Restapi0JavaApplication {
	
	public static void main(String[] args) {
		SpringApplication.run(Restapi0JavaApplication.class, args);
	}
}
