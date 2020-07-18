package api

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	automl "cloud.google.com/go/automl/apiv1"
	automlpb "google.golang.org/genproto/googleapis/cloud/automl/v1"
)

// visionClassificationPredict does a prediction for image classification.
func visionClassificationPredict(w io.Writer, projectID string, location string, modelID string, filePath string) error {
	// projectID := "my-project-id"
	// location := "us-central1"
	// modelID := "ICN123456789..."
	// filePath := "path/to/image.jpg"

	ctx := context.Background()
	client, err := automl.NewPredictionClient(ctx)
	if err != nil {
		return fmt.Errorf("NewPredictionClient: %v", err)
	}
	defer client.Close()

	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("Open: %v", err)
	}
	defer file.Close()
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return fmt.Errorf("ReadAll: %v", err)
	}

	req := &automlpb.PredictRequest{
		Name: fmt.Sprintf("projects/%s/locations/%s/models/%s", projectID, location, modelID),
		Payload: &automlpb.ExamplePayload{
			Payload: &automlpb.ExamplePayload_Image{
				Image: &automlpb.Image{
					Data: &automlpb.Image_ImageBytes{
						ImageBytes: bytes,
					},
				},
			},
		},
		// Params is additional domain-specific parameters.
		Params: map[string]string{
			// score_threshold is used to filter the result.
			"score_threshold": "0.8",
		},
	}

	resp, err := client.Predict(ctx, req)
	if err != nil {
		return fmt.Errorf("Predict: %v", err)
	}

	for _, payload := range resp.GetPayload() {
		fmt.Fprintf(w, "Predicted class name: %v\n", payload.GetDisplayName())
		fmt.Fprintf(w, "Predicted class score: %v\n", payload.GetClassification().GetScore())
	}

	return nil
}
