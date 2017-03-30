package restapi0_java;

import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.ResponseBody;

@Controller
@RequestMapping("/")
public class HomePgController {
	
	@RequestMapping(method=RequestMethod.GET)
    public @ResponseBody RandomNum sayHello(@RequestParam(value="name", required=false, defaultValue="Stranger") String name) {
        return new RandomNum();
    }

}
