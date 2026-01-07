package com.anb.admin.api;

import java.util.Optional;
import java.util.List;
import java.util.Map;
import java.util.HashMap;

import java.io.InputStreamReader;
import java.net.URL;
import java.net.URLConnection;
import java.io.BufferedReader;

import lombok.extern.slf4j.Slf4j;

import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.PutMapping;
import org.springframework.web.bind.annotation.DeleteMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RestController;
import org.springframework.data.domain.Pageable;
import org.springframework.data.domain.Page;

import org.apache.commons.lang3.StringUtils;

import com.anb.admin.domain.Report;
import com.anb.admin.domain.ReportSpecs.SearchKey;
import com.anb.admin.service.ReportService;

@Slf4j
@RestController
@RequestMapping("/api/report")
public class ReportController {

    @Autowired
    ReportService service;

    @PostMapping("")
    public ResponseEntity<? extends BasicResponse> insert(@RequestBody Report item) {
        Report result = service.insert(item);

        if (result == null) {
            return ResponseEntity.ok().body(new ErrorResponse("이미 등록된 데이터입니다", HttpStatus.CONFLICT));
        }

        return ResponseEntity.ok().body(new CommonResponse<Report>(result));
    }

    @PutMapping("{id}")
    public ResponseEntity<? extends BasicResponse> update(@PathVariable Long id, @RequestBody Report item) {
        Report result = service.update(item);

        if (result == null) {
            return ResponseEntity.ok().body(new ErrorResponse("이미 등록된 데이터입니다", HttpStatus.CONFLICT));
        }

        return ResponseEntity.ok().body(new CommonResponse<Report>(result));
    }

    @DeleteMapping("{id}")
    public ResponseEntity<? extends BasicResponse> delete(@PathVariable Long id) {
        service.delete(id);

        return ResponseEntity.noContent().build();
    }

    @GetMapping("{id}")
    public ResponseEntity<? extends BasicResponse> findById(@PathVariable Long id) {
        Optional<Report> opt = service.findById(id);

        if (!opt.isPresent()) {
            return ResponseEntity.ok().body(new ErrorResponse("데이터를 찾을수가 없습니다", HttpStatus.NOT_FOUND));
        }

        return ResponseEntity.ok().body(new CommonResponse<Report>(opt.get()));
    }

    @GetMapping("")
    public ResponseEntity<? extends BasicResponse> findAll(@RequestParam(value = "page", defaultValue = "0", required = false) int page,
                                                           @RequestParam(value = "size", defaultValue = "0", required = false) int size,
                                                           @RequestParam(value = "status", defaultValue = "0", required = false) int status) {
        Map<SearchKey, Object> searchKeys = new HashMap<>();
        if (status > 0) {
            searchKeys.put(SearchKey.valueOf("STATUS"), status);
        }

        if (size == 0) {
            List<Report> result = service.findAll(searchKeys);
            return ResponseEntity.ok().body(new CommonResponse<List<Report>>(result));
        } else {
            Page<Report> result = service.findAll(searchKeys, page, size);
            return ResponseEntity.ok().body(new CommonResponse<Page<Report>>(result));
        }
    }

    @GetMapping("summary/{id}")
    public String summary(@PathVariable Long id) {
        log.debug("Id = " + id);
        
        String urlPath = "http://localhost:3001/api/report?id=" + id;
        String pageContents = "";
        StringBuilder contents = new StringBuilder();
 
        try{
 
            URL url = new URL(urlPath);
            URLConnection con = (URLConnection)url.openConnection();
            InputStreamReader reader = new InputStreamReader (con.getInputStream(), "utf-8");
 
            BufferedReader buff = new BufferedReader(reader);
 
            while((pageContents = buff.readLine())!=null){
                //System.out.println(pageContents);             
                contents.append(pageContents);
                contents.append("\r\n");
            }
 
            buff.close();
 
            log.debug(contents.toString());
 
        }catch(Exception e){
            log.debug("error");
            e.printStackTrace();
        }
        

        return contents.toString();
    }

}
