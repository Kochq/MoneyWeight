import {
    View,
    TextInput,
    Button,
    Text,
    Pressable,
    StyleSheet,
} from "react-native";
import Screen from "./Screen";
import { useEffect, useState } from "react";
import { Picker } from "@react-native-picker/picker";
import { Category } from "@/types";
import { Ionicons } from "@expo/vector-icons";
import { useTheme } from "../../../theme/ThemeContext";
import { DateTimePickerAndroid } from "@react-native-community/datetimepicker";

export default function AddMoney() {
    const { colors } = useTheme();
    const [name, setName] = useState<string>("");
    const [money, setMoney] = useState("");
    const [categories, setCategories] = useState<Category[]>([]);
    const [selectedCategory, setSelectedCategory] = useState<Category>();
    const [date, setDate] = useState<Date>(new Date());

    const styles = StyleSheet.create({
        container: {
            justifyContent: "center",
            alignItems: "stretch",
            width: "100%",
            borderRadius: 10,
            borderWidth: 1,
            padding: 10,
            borderColor: colors.textPrimary,
            backgroundColor: colors.surface,
            boxShadow: "0 0 10px 0 rgba(0, 0, 0, 0.1)",
        },

        dateBtn: {
            backgroundColor: colors.primary,
            alignItems: "center",
            padding: 5,
            borderRadius: 5,
            margin: 5,
        },

        date: {
            flexDirection: "row",
            alignItems: "center",
            width: "100%",
            justifyContent: "center",
        },
    });

    const onChange = (event: any, selectedDate: Date) => {
        const currentDate = selectedDate;
        setDate(currentDate);
    };
    const showMode = (currentMode: any) => {
        DateTimePickerAndroid.open({
            value: date,
            onChange,
            mode: currentMode,
            is24Hour: true,
        });
    };

    const showDatepicker = () => {
        showMode("date");
    };

    const showTimepicker = () => {
        showMode("time");
    };

    const addTransaction = async () => {
        let url = "http://serkq.org:8080/api/transactions";

        console.log("Name: " + name);
        console.log("Money: " + money);
        console.log("Category: " + selectedCategory);

        // formatDate: "2023-12-27 20:18:43",
        const formatDate = `${date.getFullYear()}-${date.getMonth() + 1}-${date.getDate()} ${date.getHours()}:${date.getMinutes()}:${date.getSeconds()}`;

        try {
            const response = await fetch(url, {
                method: "POST",
                body: JSON.stringify({
                    title: name,
                    amount: Number(money),
                    category_id: 1,
                    subcategory_id: 1,
                    currency: "ARS",
                    payment_method: "cosa",
                    date: formatDate,
                    exchange_rate: 1500,
                    notes: "esto es una nota",
                    from_account_id: 1,
                }),
            });

            const data = await response.json();

            console.log("Transaction added!");
            console.log(data);
            alert(data.status);
        } catch (e) {
            console.error(e);
        }
    };

    const getCategories = async () => {
        let url = "http://serkq.org:8080/api/categories";

        const response = await fetch(url);
        const data = await response.json();

        console.log("Data fetched!");
        setCategories(data.data);
    };

    useEffect(() => {
        console.log("");
        console.log("Fetching data...");
        getCategories();
    }, []);

    return (
        <Screen>
            <View style={styles.container}>
                <TextInput
                    placeholder="Name"
                    value={name}
                    onChangeText={setName}
                />
                <TextInput
                    placeholder="Money"
                    keyboardType="numeric"
                    value={money}
                    onChangeText={setMoney}
                />
                <Picker
                    style={{ borderWidth: 1, color: colors.textPrimary }}
                    selectedValue={selectedCategory}
                    onValueChange={(itemValue, itemIndex) =>
                        setSelectedCategory(itemValue)
                    }
                >
                    {categories.map((category) => (
                        <Picker.Item
                            key={category.id}
                            label={category.name}
                            value={category.id}
                        />
                    ))}
                </Picker>
                <View style={styles.date}>
                    <Text>Date: {date.toLocaleString()}</Text>
                    <Pressable style={styles.dateBtn} onPress={showDatepicker}>
                        <Ionicons
                            name="calendar"
                            size={24}
                            color={colors.surface}
                        />
                    </Pressable>
                    <Pressable style={styles.dateBtn} onPress={showTimepicker}>
                        <Ionicons
                            name="time"
                            size={24}
                            color={colors.surface}
                        />
                    </Pressable>
                </View>
                <Pressable style={styles.dateBtn} onPress={addTransaction}>
                    <Text style={{ color: colors.surface }}>Add</Text>
                </Pressable>
            </View>
        </Screen>
    );
}
