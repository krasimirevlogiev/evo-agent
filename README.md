# Financial Agent

This is an AI-powered agent that automates the process of financial analysis and reporting. The agent reads data from Google Sheets, uses Google's Gemini AI to generate a summary, creates a new Google Sheet with the analysis, and emails a link to the report to a specified recipient.

## Features

-   **Automated Data Collection**: Automatically reads data from specified Google Sheets.
-   **AI-Powered Analysis**: Leverages the Gemini 1.5 Flash model to generate insightful financial summaries.
-   **Automated Reporting**: Creates a new Google Sheet with the generated financial summary.
-   **Email Notifications**: Sends an email notification with a link to the report.
-   **Secure Authentication**: Uses OAuth 2.0 to securely access Google services.

## How It Works

1.  **Authentication**: The agent authenticates with Google services using OAuth 2.0.
2.  **Data Extraction**: It reads data from the Google Sheets documents you specify.
3.  **AI Summarization**: The collected data is sent to the Gemini API, which returns a financial summary.
4.  **Report Generation**: A new Google Sheet is created and populated with the summary.
5.  **Notification**: An email is sent via Gmail with a link to the newly created report.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

-   Go (version 1.18 or later)
-   A Google Cloud Platform project
-   Go modules enabled

### Installation & Configuration

1.  **Clone the repository:**
    ```sh
    git clone <repository-url>
    cd financial-agent
    ```

2.  **Enable Google APIs:**

    Enable the Google Sheets API and Gmail API in the [Google Cloud Console](https://console.cloud.google.com/).

3.  **Create OAuth 2.0 Credentials:**

    Create OAuth 2.0 client ID credentials and download the `credentials.json` file. Place it in the root of the project directory.

4.  **Set up Environment Variables:**

    Create a `.env` file in the root of the project and add your Gemini API key:
    ```
    GEMINI_API_KEY=your_gemini_api_key
    ```

5.  **Configure the Agent:**

    Open `cmd/financial-agent/main.go` and update the following variables:
    - `spreadsheetIDs`: A slice of strings containing the IDs of the Google Sheets you want to analyze.
    - `recipientEmail`: The email address where the report will be sent.

6.  **Install dependencies:**
    ```sh
    go mod tidy
    ```

## Usage

1.  **First-time Authentication:**

    When you run the agent for the first time, it will prompt you to visit a URL to authorize the application. After authorization, you will receive a code to paste back into the terminal. This will create a `token.json` file for future authentications.

2.  **Run the agent:**
    ```sh
    make run
    ```
    Alternatively, you can run the `main.go` file directly:
    ```sh
    go run cmd/financial-agent/main.go
    ```
3.  **Build the agent:**
    To build a production executable:
    ```sh
    make build
    ```

## Future Improvements

-   Move configuration (spreadsheet IDs, email) to a separate config file.
-   Error handling and retry mechanisms.
-   Support for more data sources (e.g., CSV files, databases).
-   More advanced financial analysis prompts.

---
