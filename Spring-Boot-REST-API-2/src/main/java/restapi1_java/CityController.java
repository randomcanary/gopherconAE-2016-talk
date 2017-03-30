package restapi1_java;

import java.util.ArrayList;
import java.util.List;
import java.util.stream.Collectors;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.jdbc.core.JdbcTemplate;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.ResponseBody;


@Controller
public class CityController {

    @Autowired
    private JdbcTemplate jdbcTemplate;

    @RequestMapping(value = "/cities")
    public
    @ResponseBody
    List<City> getSortedCities() {
        String query = "SELECT Name from City" ;
    	List<String> citiez = (List<String>) jdbcTemplate.queryForList(query, String.class);
        List<String> sortedCitiez = citiez.stream().sorted().collect(Collectors.toList());
        
        List<City> list_cities = new ArrayList<City>();
        sortedCitiez.forEach((element) -> {
        	City a = new City(element);
        	list_cities.add(a);
		});

		return list_cities;

    }
}