package com.anb.admin;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.WebApplicationType;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.boot.builder.SpringApplicationBuilder;
import org.springframework.boot.context.ApplicationPidFileWriter;
import org.springframework.boot.web.servlet.ServletComponentScan;

@ServletComponentScan
@SpringBootApplication
public class BootSpringBootApplication {
	public static void main(String[] args) {
		// SpringApplication.run(BootSpringBootApplication.class, args);

		SpringApplication app = new SpringApplication(BootSpringBootApplication.class);
		app.addListeners(new ApplicationPidFileWriter());
		app.setWebApplicationType(WebApplicationType.SERVLET);
		app.run(args);
	}
}
