package com.torttable.restfulAPI.model;

import java.math.BigDecimal;

public class Plant {
    private String id;
    private String name;
    private String safety;
    private String latinName;
    private String familyName;
    private String description;

    public Plant(final String id, final String Name, final String safety, final String latinName, final String familyName, final String description) {
        this.id = id;
        this.name = Name;
        this.safety = safety;
        this.latinName = latinName;
        this.familyName = familyName;
        this.description = description;
    }

    public String getId() {
        return id;
    }

    public void setId(String id) {
        this.id = id;
    }

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public String getSafety() {
        return safety;
    }

    public void setSafety(String safety) {
        this.safety = safety;
    }

    public String getLatinName() {
        return latinName;
    }

    public void setLatinName(String latinName) {
        this.latinName = latinName;
    }

    public String getFamilyName() {
        return familyName;
    }

    public void setFamilyName(String familyName) {
        this.familyName = familyName;
    }

    public String getDescription() {
        return description;
    }

    public void setDescription(String description) {
        this.description = description;
    }


    @Override
    public String toString() {
        return "Plant{" +
                "id='" + id + '\'' +
                ", name ='" + name + '\'' +
                ", safety=" + safety + '\'' +
                ", latinName=" + latinName + '\'' +
                ", familyName=" + familyName + '\'' +
                ", description=" + description +
                '}';
    }
}
