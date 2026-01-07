package com.anb.admin.api;

import java.util.List;

import org.springframework.http.HttpStatus;

import lombok.Getter;
import lombok.Setter;

@Getter
@Setter
public class ErrorResponse extends BasicResponse {
	private String message;
	private HttpStatus code;

	public ErrorResponse(String message) {
		this.message = message;
		this.code = HttpStatus.NOT_FOUND;
	}
	public ErrorResponse(String message, HttpStatus code) {
		this.message = message;
		this.code = code;
	}
}
