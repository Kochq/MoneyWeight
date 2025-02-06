import {
    View,
    TextInput,
    Text,
    Pressable,
    StyleSheet,
    KeyboardAvoidingView,
    Platform,
    ScrollView,
} from "react-native";
import Screen from "./Screen";
import { useEffect, useState } from "react";
import { Picker } from "@react-native-picker/picker";
import { Category, SubCategory } from "@/types";
import { Ionicons } from "@expo/vector-icons";
import { useTheme } from "../../../theme/ThemeContext";
import { DateTimePickerAndroid } from "@react-native-community/datetimepicker";

export default function AddMoney() {
    const { colors } = useTheme();
    const [name, setName] = useState<string>("");
    const [money, setMoney] = useState("");
    const [categories, setCategories] = useState<Category[]>([]);
    const [selectedCategory, setSelectedCategory] = useState<Category>();
    const [selectedSubCategory, setSelectedSubCategory] =
        useState<SubCategory>();
    const [date, setDate] = useState<Date>(new Date());
    const apiUrl = process.env.EXPO_PUBLIC_API_BASE;

    if (selectedCategory) {
        console.log(selectedCategory.subcategories);
    }

    const styles = StyleSheet.create({
        container: {
            flex: 1,
            backgroundColor: colors.background,
        },

        card: {
            backgroundColor: colors.surface,
            minWidth: "100%",
            borderRadius: 15,
            padding: 20,
            boxShadow: "0 0 10px rgba(0, 0, 0, 0.1)",
        },

        inputContainer: {
            marginBottom: 16,
        },

        label: {
            fontSize: 14,
            color: colors.textSecondary,
            marginBottom: 8,
            fontWeight: "500",
        },

        input: {
            backgroundColor: colors.background,
            borderRadius: 10,
            padding: 12,
            fontSize: 16,
            color: colors.textPrimary,
            borderWidth: 1,
            borderColor: colors.border,
        },

        picker: {
            backgroundColor: colors.background,
            borderRadius: 10,
            marginBottom: 16,
            borderWidth: 1,
            borderColor: colors.border,
        },

        dateContainer: {
            flexDirection: "row",
            alignItems: "center",
            marginBottom: 20,
            backgroundColor: colors.background,
            borderRadius: 10,
            padding: 12,
            borderWidth: 1,
            borderColor: colors.border,
        },

        dateText: {
            flex: 1,
            color: colors.textPrimary,
            fontSize: 16,
        },

        dateButton: {
            padding: 8,
            borderRadius: 8,
            marginLeft: 8,
            backgroundColor: colors.primary + "20",
        },

        addButton: {
            backgroundColor: colors.primary,
            borderRadius: 12,
            padding: 16,
            alignItems: "center",
            marginTop: 10,
        },

        addButtonText: {
            color: colors.surface,
            fontSize: 16,
            fontWeight: "600",
        },

        moneyInput: {
            fontSize: 24,
            textAlign: "center",
            color: colors.primary,
            fontWeight: "bold",
        },

        titleInput: {
            textAlign: "center",
            color: colors.textPrimary,
        },
    });

    const onChange = (_event: any, selectedDate: Date) => {
        const currentDate = selectedDate;
        setDate(currentDate);
    };

    const showMode = (currentMode: any) => {
        DateTimePickerAndroid.open({
            value: date,
            // @ts-ignore
            onChange,
            mode: currentMode,
            is24Hour: true,
        });
    };

    const showDatepicker = () => showMode("date");
    const showTimepicker = () => showMode("time");

    const addTransaction = async () => {
        if (!name || !money) {
            alert("Por favor completa todos los campos");
            return;
        }

        const formatDate = `${date.getFullYear()}-${date.getMonth() + 1}-${date.getDate()} ${date.getHours()}:${date.getMinutes()}:${date.getSeconds()}`;

        const url = apiUrl + "/api/transactions";

        try {
            const response = await fetch(url, {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify({
                    title: name,
                    amount: Number(money),
                    category_id: selectedCategory?.id || 1,
                    subcategory_id: selectedSubCategory?.id || 1,
                    currency: "ARS",
                    payment_method: "cosa",
                    date: formatDate,
                    exchange_rate: 1500,
                    notes: "esto es una nota",
                    from_account_id: 1,
                }),
            });

            const data = await response.json();
            alert(data.status);

            // Limpiar formulario
            setName("");
            setMoney("");
            setDate(new Date());
        } catch (e) {
            alert("Error al guardar la transacción");
            console.error(e);
        }
    };

    useEffect(() => {
        const getCategories = async () => {
            try {
                const url = apiUrl + "/api/categories";
                console.log(url);
                const response = await fetch(url);
                const data = await response.json();
                setCategories(data.data);
            } catch (e) {
                console.error("Error al obtener categorías:", e);
            }
        };

        getCategories();
    }, [apiUrl]);

    return (
        <Screen>
            <KeyboardAvoidingView
                behavior={Platform.OS === "ios" ? "padding" : "height"}
                style={styles.container}
            >
                <ScrollView>
                    <View style={styles.card}>
                        <View style={styles.inputContainer}>
                            <Text style={styles.label}>Amount</Text>
                            <TextInput
                                style={[styles.input, styles.moneyInput]}
                                placeholder="$0.00"
                                keyboardType="numeric"
                                value={money}
                                placeholderTextColor={colors.textSecondary}
                                onChangeText={setMoney}
                            />
                        </View>

                        <View style={styles.inputContainer}>
                            <Text style={styles.label}>Title</Text>
                            <TextInput
                                style={[styles.input, styles.titleInput]}
                                placeholder="What did you buy?"
                                value={name}
                                placeholderTextColor={colors.textSecondary}
                                onChangeText={setName}
                            />
                        </View>

                        <Text style={styles.label}>Category</Text>
                        <View style={styles.picker}>
                            <Picker
                                selectedValue={selectedCategory}
                                onValueChange={setSelectedCategory}
                                style={{ color: colors.textPrimary }}
                                prompt="Select a category"
                            >
                                <Picker.Item
                                    label="Select category"
                                    value={null}
                                />
                                {categories.map((category) => (
                                    <Picker.Item
                                        key={category.id}
                                        label={
                                            category.icon + " " + category.name
                                        }
                                        value={category}
                                    />
                                ))}
                            </Picker>
                        </View>

                        {selectedCategory &&
                            selectedCategory.subcategories.length > 0 && (
                                <>
                                    <Text style={styles.label}>
                                        SubCategory
                                    </Text>
                                    <View style={styles.picker}>
                                        <Picker
                                            selectedValue={selectedSubCategory}
                                            onValueChange={
                                                setSelectedSubCategory
                                            }
                                            style={{
                                                color: colors.textPrimary,
                                            }}
                                        >
                                            <Picker.Item
                                                label="Select subCategory"
                                                value={null}
                                            />
                                            {selectedCategory?.subcategories.map(
                                                (subCategory) => (
                                                    <Picker.Item
                                                        key={subCategory.id}
                                                        label={
                                                            subCategory.icon +
                                                            " " +
                                                            subCategory.name
                                                        }
                                                        value={subCategory}
                                                    />
                                                ),
                                            )}
                                        </Picker>
                                    </View>
                                </>
                            )}

                        <View style={styles.inputContainer}>
                            <Text style={styles.label}>Date</Text>
                            <View style={styles.dateContainer}>
                                <Text style={styles.dateText}>
                                    {date.toLocaleString()}
                                </Text>
                                <Pressable
                                    style={styles.dateButton}
                                    onPress={showDatepicker}
                                >
                                    <Ionicons
                                        name="calendar"
                                        size={24}
                                        color={colors.primary}
                                    />
                                </Pressable>
                                <Pressable
                                    style={styles.dateButton}
                                    onPress={showTimepicker}
                                >
                                    <Ionicons
                                        name="time"
                                        size={24}
                                        color={colors.primary}
                                    />
                                </Pressable>
                            </View>
                        </View>

                        <Pressable
                            style={styles.addButton}
                            onPress={addTransaction}
                        >
                            <Text style={styles.addButtonText}>
                                Agregar Transacción
                            </Text>
                        </Pressable>
                    </View>
                </ScrollView>
            </KeyboardAvoidingView>
        </Screen>
    );
}
