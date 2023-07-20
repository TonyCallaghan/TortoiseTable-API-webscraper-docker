package com.torttable.restfulAPI.contoller;

import com.torttable.restfulAPI.contoller.model.PlantID;
import org.springframework.http.HttpStatus;
import org.springframework.http.MediaType;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;
import com.torttable.restfulAPI.model.Plant;

import java.util.*;
import java.util.concurrent.ConcurrentHashMap;

@RestController
public class PlantController {
    private Map<String, Plant> plantMap = new ConcurrentHashMap<>();

    @PostMapping("/plant")
    public ResponseEntity<PlantID> createPlant(@RequestBody final Plant plant) {
        PlantID result = new PlantID(UUID.randomUUID().toString());
        plant.setId(result.getId());
        plantMap.put(result.getId(), plant);
        return ResponseEntity.status(HttpStatus.CREATED)
                .contentType(MediaType.APPLICATION_JSON)
                .body(result);
    }

    @GetMapping("/plant")
    public List<Plant> getAllPlants() {
        if(plantMap.isEmpty()) {
            populatePlants();
        }
        return new ArrayList<>(plantMap.values());
    }

    @GetMapping("/plant/{id}")
    public ResponseEntity<Plant> getPlantById(@PathVariable final String id) {
        Plant plant = plantMap.get(id);
        if (plant == null) {
            return ResponseEntity.status(HttpStatus.NOT_FOUND).body(null);
        } else {
            return ResponseEntity.ok(plant);
        }
    }

    @PatchMapping("/plant/{id}")
    public ResponseEntity<Plant> patchPlant(@PathVariable final String id, @RequestBody Plant plantUpdates) {
        Plant plant = plantMap.get(id);
        if (plant == null) {
            return ResponseEntity.status(HttpStatus.NOT_FOUND).body(null);
        } else {
            plant.setName(plantUpdates.getName());
            plant.setSafety(plantUpdates.getSafety());
            plant.setLatinName(plantUpdates.getLatinName());
            plant.setFamilyName(plantUpdates.getFamilyName());
            plant.setDescription(plantUpdates.getDescription());
            return ResponseEntity.ok(plant);
        }
    }

    @DeleteMapping("/plant/{id}")
    public ResponseEntity<String> deletePlant(@PathVariable final String id) {
        Plant removedPlant = plantMap.remove(id);
        if (removedPlant == null) {
            return ResponseEntity.status(HttpStatus.NOT_FOUND).body("Plant not found");
        } else {
            return ResponseEntity.ok("Plant deleted successfully");
        }
    }

    private void populatePlants() { // Sample data
        Plant plant1 = new Plant("1",
                "Abutilon (Flowering Maple, Chinese Lantern)",
                "Safe to Feed",
                "Abutilon spp.",
                "Malvaceae",
                "No known toxicity to animals although occasionally may cause skin irritation and allergy in humans. " +
                "Tortoises love the sweet flowers when they are available." +
                "Often grown as a pot plant in the UK as this plant is sensitive to frosts."
        );
        Plant plant2 = new Plant("2",
                "Aconite (Monkshood, Wolfsbane)",
                "Do not Feed",
                "Aconitum spp.",
                "Ranunculaceae",
                "This plant is toxic. All parts of this plant contain the highly toxic alkaloid aconitine which initially" +
                " paralyses the nervous system before lowering pulse rate and stopping the heart in animals. " +
                "Poisoning is not only via ingestion but contact through unbroken skin; so if clearing away from the garden care should be taken."
        );

        plantMap.put(plant1.getId(), plant1);
        plantMap.put(plant2.getId(), plant2);
    }
}
