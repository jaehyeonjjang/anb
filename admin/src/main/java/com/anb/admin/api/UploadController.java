package com.anb.admin.api;

import java.util.Optional;
import java.util.List;
import java.util.UUID;
import java.util.Date;
import java.text.SimpleDateFormat;
import java.io.File;
import java.io.UnsupportedEncodingException;
import java.nio.file.Files;
import java.nio.file.Path;
import java.nio.file.Paths;

import lombok.extern.slf4j.Slf4j;

import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.http.HttpHeaders;

import org.springframework.core.io.Resource;
import org.springframework.core.io.FileSystemResource;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.PutMapping;
import org.springframework.web.bind.annotation.DeleteMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RestController;
import org.springframework.data.domain.Pageable;
import org.springframework.data.domain.Page;

import org.springframework.web.multipart.MultipartFile;
import org.springframework.http.ContentDisposition;

import java.nio.charset.StandardCharsets;

import com.anb.admin.domain.Apt;
import com.anb.admin.service.AptService;
import com.anb.admin.domain.Aptgroup;
import com.anb.admin.service.AptgroupService;

@Slf4j
@RestController
public class UploadController {

    @Value("${path.root-path}")
    private String rootPath;

    @Value("${path.upload-path}")
    private String uploadPath;

    @Autowired
    AptService service;

    @Autowired
    AptgroupService aptgroupService;

    @RequestMapping(value = "/api/upload", method = RequestMethod.POST)
    public ResponseEntity<? extends BasicResponse> upload(@RequestParam("file") MultipartFile file) {
        String filename = "";
        if (!file.isEmpty()) {
            log.info("file upload");

            try {
                int begin = file.getOriginalFilename().indexOf(".");
                int last = file.getOriginalFilename().length();
                String ext = file.getOriginalFilename().substring(begin, last);
                filename = UUID.randomUUID().toString().replace("-", "").toLowerCase() + ext;
                String fullPath = rootPath + "/" + uploadPath + "/original";

                File file1 = new File(fullPath);
                if (!file1.exists()) {
                    file1.mkdirs();
                }

                String fullFilename = fullPath + "/" + filename;

                file.transferTo(new File(fullFilename));
            } catch (Exception e) {
                e.printStackTrace();
            }
        } else {
            log.info("file empty");
            return ResponseEntity.status(HttpStatus.NOT_FOUND).body(new ErrorResponse("데이터를 찾을수가 없습니다"));
        }

        return ResponseEntity.ok().body(new CommonResponse<String>(uploadPath + "/" + filename));
    }

    @GetMapping(value = "/api/download")
    public ResponseEntity<Resource> download(@RequestParam(value = "filename") String filename) {
        String fullFilename = rootPath + "/" + filename.replace("webdata/", "webdata/original/");
        Resource resource = new FileSystemResource(fullFilename);

        log.info("download image : " + fullFilename);

        if (!resource.exists()) {
            return new ResponseEntity<Resource>(HttpStatus.NOT_FOUND);
        }

        HttpHeaders header = new HttpHeaders();

        try {
            Path filePath = Paths.get(fullFilename);
            header.add("Content-Type", Files.probeContentType(filePath));
        } catch (Exception e) {
            e.printStackTrace();
        }

        return new ResponseEntity<Resource>(resource, header, HttpStatus.OK);
    }

    @RequestMapping(value = "/api/upload/management", method = RequestMethod.POST)
    public ResponseEntity<? extends BasicResponse> uploadCompany(@RequestParam("file") MultipartFile file) {
        String filename = "";
        if (!file.isEmpty()) {
            log.info("file upload");

            try {
                int begin = file.getOriginalFilename().indexOf(".");
                int last = file.getOriginalFilename().length();
                String ext = file.getOriginalFilename().substring(begin, last);
                filename = UUID.randomUUID().toString().replace("-", "").toLowerCase() + ext;
                String fullPath = rootPath + "/" + uploadPath + "/management";

                File file1 = new File(fullPath);
                if (!file1.exists()) {
                    file1.mkdirs();
                }

                String fullFilename = fullPath + "/" + filename;

                file.transferTo(new File(fullFilename));
            } catch (Exception e) {
                e.printStackTrace();
            }
        } else {
            log.info("file empty");
            return ResponseEntity.status(HttpStatus.NOT_FOUND).body(new ErrorResponse("데이터를 찾을수가 없습니다"));
        }

        return ResponseEntity.ok().body(new CommonResponse<String>(uploadPath + "/management/" + filename));
    }

    @GetMapping(value = "/api/download/management")
    public ResponseEntity<Resource> downloadManagement(@RequestParam(value = "filename") String filename) {
        String fullFilename = rootPath + "/" + filename.replace("../", "");
        Resource resource = new FileSystemResource(fullFilename);

        log.info("download image : " + fullFilename);

        if (!resource.exists()) {
            return new ResponseEntity<Resource>(HttpStatus.NOT_FOUND);
        }

        HttpHeaders header = new HttpHeaders();

        try {
            Path filePath = Paths.get(fullFilename);
            header.add("Content-Type", Files.probeContentType(filePath));
        } catch (Exception e) {
            e.printStackTrace();
        }

        return new ResponseEntity<Resource>(resource, header, HttpStatus.OK);
    }

    @GetMapping(value = "/api/download/report/{id}")
    public ResponseEntity<Resource> downloadReport(@PathVariable Long id) {
        Optional<Apt> opt = service.findById(id);

        if (!opt.isPresent()) {
            return new ResponseEntity<Resource>(null, null, HttpStatus.OK);
        }

        Apt apt = opt.get();

        Optional<Aptgroup> optAptgroup = aptgroupService.findById(apt.getAptgroup());

        if (!optAptgroup.isPresent()) {
            return new ResponseEntity<Resource>(null, null, HttpStatus.OK);
        }

        Aptgroup aptgroup = optAptgroup.get();

        String title = aptgroup.getName() + " " + apt.getName() + ".zip";

        String fullFilename = rootPath + "/" + uploadPath + "/" + id + ".zip";
        Resource resource = new FileSystemResource(fullFilename);

        log.info("download report : " + fullFilename);

        if (!resource.exists()) {
            return new ResponseEntity<Resource>(HttpStatus.NOT_FOUND);
        }

        HttpHeaders header = new HttpHeaders();

        try {
            Path filePath = Paths.get(fullFilename);
            header.add("Content-Type", Files.probeContentType(filePath));
            header.setContentDisposition(ContentDisposition.builder("attachment").filename(title, StandardCharsets.UTF_8).build());
        } catch (Exception e) {
            e.printStackTrace();
        }

        return new ResponseEntity<Resource>(resource, header, HttpStatus.OK);
    }
}
